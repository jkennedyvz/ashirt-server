// Copyright 2020, Verizon Media
// Licensed under the terms of the MIT. See LICENSE file in project root for terms.

package services

import (
	"context"

	"github.com/theparanoids/ashirt/backend"
	"github.com/theparanoids/ashirt/backend/database"
	"github.com/theparanoids/ashirt/backend/policy"
	"github.com/theparanoids/ashirt/backend/server/middleware"

	sq "github.com/Masterminds/squirrel"
)

type DeleteEvidenceInput struct {
	OperationSlug            string
	EvidenceUUID             string
	DeleteAssociatedFindings bool
}

func DeleteEvidence(ctx context.Context, db *database.Connection, i DeleteEvidenceInput) error {
	operation, evidence, err := lookupOperationEvidence(db, i.OperationSlug, i.EvidenceUUID)
	if err != nil {
		return backend.UnauthorizedWriteErr(err)
	}

	if err := policy.Require(middleware.Policy(ctx), policy.CanModifyEvidenceOfOperation{OperationID: operation.ID}); err != nil {
		return backend.UnauthorizedWriteErr(err)
	}

	err = db.WithTx(ctx, func(tx *database.Transactable) {
		if i.DeleteAssociatedFindings {
			tx.Exec(sq.Expr("DELETE findings FROM findings INNER JOIN evidence_finding_map ON findings.id = evidence_finding_map.finding_id WHERE evidence_id = ?", evidence.ID))
		}
		tx.Delete(sq.Delete("evidence_finding_map").Where(sq.Eq{"evidence_id": evidence.ID}))
		tx.Delete(sq.Delete("evidence").Where(sq.Eq{"id": evidence.ID}))
	})
	if err != nil {
		return backend.DatabaseErr(err)
	}

	return nil
}
