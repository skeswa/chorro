/** Settings and values that affect how this web application should behave. */
export const config = {
  /** Base URL of the API server. */
  apiBaseUrl: import.meta.env.VITE_API_BASE_URL as string | undefined,
} as const;
