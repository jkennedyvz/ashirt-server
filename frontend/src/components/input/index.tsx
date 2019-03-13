// Copyright 2020, Verizon Media
// Licensed under the terms of the MIT. See LICENSE file in project root for terms.

import * as React from 'react'
import LoadingSpinner from 'src/components/loading_spinner'
import WithLabel from 'src/components/with_label'
import classnames from 'classnames/bind'
const cx = classnames.bind(require('./stylesheet'))

type SharedProps = {
  className?: string,
  disabled?: boolean,
  label?: string,
  name?: string,
  onBlur?: () => void,
  onChange?: (newValue: string) => void,
  onClick?: () => void,
  onFocus?: () => void,
  onKeyDown?: (e: React.KeyboardEvent) => void,
  placeholder?: string,
  readOnly?: boolean,
  value?: string,
}

export default React.forwardRef((props: SharedProps & {
  icon?: string,
  loading?: boolean,
  type?: string,
}, ref: React.RefObject<HTMLInputElement>) => (
  <WithLabel className={cx('root', props.className)} label={props.label}>
    {props.loading && (
      <LoadingSpinner className={cx('spinner')} small />
    )}
    <input
      ref={ref}
      className={cx('input', {'has-icon': props.icon != null, loading: props.loading != null})}
      disabled={props.disabled}
      name={props.name}
      onBlur={props.onBlur}
      onChange={e => { if (props.onChange) props.onChange(e.target.value) }}
      onClick={props.onClick}
      onFocus={props.onFocus}
      onKeyDown={props.onKeyDown}
      placeholder={props.placeholder}
      readOnly={props.readOnly}
      style={props.icon != null ? {backgroundImage: `url(${props.icon})`} : {}}
      type={props.type}
      value={props.value}
    />
  </WithLabel>
))

export const TextArea = React.forwardRef((props: SharedProps & {
}, ref: React.RefObject<HTMLTextAreaElement>) => (
  <WithLabel className={cx('root', props.className)} label={props.label}>
    <textarea
      ref={ref}
      className={cx('input', 'textarea')}
      disabled={props.disabled}
      name={props.name}
      onBlur={props.onBlur}
      onChange={e => { if (props.onChange) props.onChange(e.target.value) }}
      onClick={props.onClick}
      onFocus={props.onFocus}
      onKeyDown={props.onKeyDown}
      placeholder={props.placeholder}
      value={props.value}
    />
  </WithLabel>
))
