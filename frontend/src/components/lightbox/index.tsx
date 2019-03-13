// Copyright 2020, Verizon Media
// Licensed under the terms of the MIT. See LICENSE file in project root for terms.

import * as React from 'react'
import classnames from 'classnames/bind'
import {createPortal} from 'react-dom'
const cx = classnames.bind(require('./stylesheet'))

export default (props: {
  children: React.ReactNode,
  isOpen: boolean,
  onRequestClose: () => void,
}) => {
  const [exists, setExists] = React.useState<boolean>(false)
  const [animating, setAnimating] = React.useState<boolean>(true)


  React.useEffect(() => {
    if (props.isOpen) {
      setExists(true)
      setAnimating(true)
      setTimeout(() => setAnimating(false), 20)
    } else {
      setAnimating(true)
      setTimeout(() => setExists(false), 200)
    }
  }, [props.isOpen])

  const onKeyDown = (e: KeyboardEvent) => {
    if (e.key === 'Escape') props.onRequestClose()
  }

  React.useEffect(() => {
    document.addEventListener('keydown', onKeyDown)
    return () => document.removeEventListener('keydown', onKeyDown)
  })

  if (!exists) return null
  return (
    createPortal((
      <div className={cx('root', animating ? 'animating' : 'open')} onClick={props.onRequestClose}>
        <div className={cx('content')} onClick={e => e.stopPropagation()}>
          {props.children}
        </div>
      </div>
    ), document.body)
  )
}
