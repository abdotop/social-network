import type { Ref } from 'vue'

export const usePersistedState = <T extends Object | null>(identifier: string, defaultOptions: T): Ref<T> => {
  const persistedObject = useState<T>(identifier, (): T => {
    const item = localStorage.getItem(identifier)
    if (!item)
      return defaultOptions

    return JSON.parse(item) as T ?? defaultOptions
  })

  watch(persistedObject, (object) => {
    localStorage.setItem(identifier, JSON.stringify(object))
  }, { deep: true })

  return persistedObject
}