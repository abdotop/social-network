import type { UserWithoutPassword } from "~/types";

export const useAuthUser = () => {
  const currentUser = usePersistedState<UserWithoutPassword | null>("user", null);
  return currentUser;
};