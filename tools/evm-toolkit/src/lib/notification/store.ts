import type { INotification } from './interfaces'
import { NOTIFICATION_TIMEOUT_DEFAULT, NOTIFICATION_TIMEOUT_NEVER } from './constants'
import { writable } from 'svelte/store'

export const notificationStore = writable<INotification[]>([])

export function showNotification(notification: Omit<INotification, 'id'>): string {
    const id = crypto.randomUUID()
    const timeout = notification.timeout ?? NOTIFICATION_TIMEOUT_DEFAULT

    notificationStore.update((notifications) => [...notifications, { ...notification, id, timeout }])

    if (timeout !== NOTIFICATION_TIMEOUT_NEVER) {
        setTimeout(() => removeNotification(id), timeout ?? NOTIFICATION_TIMEOUT_DEFAULT)
    }

    return id
}

export function removeNotification(id: string): void {
    notificationStore.update((notifications) => notifications.filter((n) => n.id !== id))
}

export function updateNotification(id: string, update: Partial<INotification>): void {
    notificationStore.update((notifications) => {
        const notification = notifications.find((n) => n.id === id)
        if (notification) {
            Object.assign(notification, update)
            if (notification.timeout !== NOTIFICATION_TIMEOUT_NEVER) {
                setTimeout(() => removeNotification(id), notification.timeout ?? NOTIFICATION_TIMEOUT_DEFAULT)
            }
        }
        return notifications
    })
}
