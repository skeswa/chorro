import type { RequestHandler } from '@sveltejs/kit';
import { config } from '$lib/config';
import type { Locals } from '$lib/types';

// GET /debug/hello.text
export const get: RequestHandler<Locals> = async () => {
  const { serverBaseUrl } = config;
  if (!serverBaseUrl) {
    return { body: 'oops' };
  }

  try {
    const response = await fetch(`${serverBaseUrl}/api`);

    if (!response.ok) {
      return { body: 'oh no!' };
    }

    return {
      body: await response.text(),
    };
  } catch (e) {
    console.error(e);

    return { body: 'rats' };
  }
};
