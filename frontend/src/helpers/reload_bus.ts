// Copyright 2020, Verizon Media
// Licensed under the terms of the MIT. See LICENSE file in project root for terms.

import { EventEmitter } from 'events'

export const reloadEvent = "reload"
export const reloadDoneEvent = "reload-done"

export const BuildReloadBus = () => {
  const bus = new EventEmitter()

  return {
    requestReload: () => {
      bus.emit(reloadEvent, null)
    },
    onReload: (listener: () => void) => {
      bus.on(reloadEvent, listener)
    },
    offReload: (listener: () => void) => {
      bus.removeListener(reloadEvent, listener)
    },

    reloadDone: () => {
      bus.emit(reloadDoneEvent, null)
    },
    onReloadDone: (listener: () => void) => {
      bus.on(reloadDoneEvent, listener)
    },
    offReloadDone: (listener: () => void) => {
      bus.removeListener(reloadDoneEvent, listener)
    },

    clean: () => {
      bus.removeAllListeners()
    }
  }
}
