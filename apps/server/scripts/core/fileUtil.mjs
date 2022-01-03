import { dirname, join as asPath } from 'path';
import { fileURLToPath } from 'url';

/** @returns the path to the root of this project */
export function resolveProjectRootDirectoryPath() {
  const scriptFilePath = fileURLToPath(import.meta.url);

  return asPath(dirname(scriptFilePath), '..', '..', '..', '..');
}
