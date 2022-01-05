/** Settings and values that affect how this web application should behave. */
export const config = {
  /** Base URL of the server. */
  serverBaseUrl: import.meta.env.VITE_SERVER_BASE_URL as string | undefined,
} as const;
