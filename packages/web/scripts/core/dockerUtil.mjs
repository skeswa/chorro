import { readFileSync } from 'fs';
import { dirname, join as asPath } from 'path';

import {
  containerRegistryDomain,
  webDockerImageName,
  webDockerImageOwner,
} from './dockerConstants.mjs';
import { packageJsonFileName } from './fileConstants.mjs';

/**
 * Current tag that should be applied to the `web` docker image.
 *
 * @param {*} packageDirectoryPath path to the `web` package directory
 */
export function composeWebDockerImageTag(packageDirectoryPath) {
  const packageJsonFilePath = asPath(packageDirectoryPath, packageJsonFileName);
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

    return (
      `${containerRegistryDomain}/${webDockerImageOwner}/` +
      `${webDockerImageName}:${version}`
    );
  } catch (err) {
    console.error(`Failed to read "${packageJsonFilePath}":`, err);

    process.exit(1);
  }
}
