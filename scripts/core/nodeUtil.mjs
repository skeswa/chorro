import { readFileSync } from 'fs';
import { join as asPath } from 'path';

import { packageJsonFileName } from './fileConstants.mjs';

/**
 * Reads the `version` of the Node.js package located within
 * `nodePackageDirectoryPath`.
 */
export function readPackageVersion(nodePackageDirectoryPath) {
  const packageJsonFilePath = asPath(
    nodePackageDirectoryPath,
    packageJsonFileName,
  );
  try {
    const packageJsonFileContents = readFileSync(packageJsonFilePath, 'utf8');
    const packageJson = JSON.parse(packageJsonFileContents);

    const { version } = packageJson;
    if (
      !version ||
      (typeof version !== 'string' && !(version instanceof String))
    ) {
      throw Error(`version "${version}" is invalid`);
    }

    return version;
  } catch (err) {
    throw new Error(`Failed to read "${packageJsonFilePath}": ${err.message}`);
  }
}
