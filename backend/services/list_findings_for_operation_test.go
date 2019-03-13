// Copyright 2020, Verizon Media
// Licensed under the terms of the MIT. See LICENSE file in project root for terms.

package services_test

import (
	"testing"

	"github.com/theparanoids/ashirt/backend/dtos"
	"github.com/theparanoids/ashirt/backend/helpers"
	"github.com/theparanoids/ashirt/backend/models"
	"github.com/theparanoids/ashirt/backend/policy"
	"github.com/theparanoids/ashirt/backend/services"
	"github.com/stretchr/testify/require"
)

type findingValidator func(*testing.T, models.Finding, *dtos.Finding)

func TestListFindingsForOperation(t *testing.T) {
	db := initTest(t)
	HarryPotterSeedData.ApplyTo(t, db)
	ctx := fullContext(UserRon.ID, &policy.FullAccess{})

	masterOp := OpChamberOfSecrets
	input := services.ListFindingsForOperationInput{
		OperationSlug: masterOp.Slug,
		Filters:       helpers.TimelineFilters{},
	}

	allFindings := getFindingsByOperationID(t, db, masterOp.ID)
	require.NotEqual(t, len(allFindings), 0, "Some number of findings should exist")

	foundFindings, err := services.ListFindingsForOperation(ctx, db, input)
	require.NoError(t, err)
	require.Equal(t, len(foundFindings), len(allFindings))
	validateFindingSets(t, foundFindings, allFindings, validateFinding)
}

func validateFinding(t *testing.T, expected models.Finding, actual *dtos.Finding) {
	require.Equal(t, expected.UUID, actual.UUID)
	require.Equal(t, expected.Category, actual.Category)
	require.Equal(t, expected.Title, actual.Title)
	require.Equal(t, expected.Description, actual.Description)
	require.Equal(t, expected.ReadyToReport, actual.ReadyToReport)
	require.Equal(t, expected.TicketLink, actual.TicketLink)
}

func validateFindingSets(t *testing.T, dtoSet []*dtos.Finding, dbSet []models.Finding, validate findingValidator) {
	var expected *models.Finding = nil

	for _, dtoItem := range dtoSet {
		expected = nil
		for _, dbItem := range dbSet {
			if dbItem.UUID == dtoItem.UUID {
				expected = &dbItem
				break
			}
		}
		require.NotNil(t, expected, "Result should have matching value")
		validate(t, *expected, dtoItem)
	}
}
