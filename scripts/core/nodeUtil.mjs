import { readFileSync } from 'fs';
import { join as asPath } from 'path';

import { packageJsonFileName } from './fileConstants.mjs';

/**
 * Reads metadata of the Node.js package located within
 * `nodePackageDirectoryPath`.
 */
export function readPackageMetadata(nodePackageDirectoryPath) {
  const packageJsonFilePath = asPath(
    nodePackageDirectoryPath,
    packageJsonFileName,
  );
  try {
    const packageJsonFileContents = readFileSync(packageJsonFilePath, 'utf8');
    const packageJson = JSON.parse(packageJsonFileContents);

    const name = readStringField({
      fieldName: 'name',
      packageJson,
    });
    const version = readStringField({
      fieldName: 'version',
      packageJson,
    });

    return { name, version };
  } catch (err) {
    throw new Error(`Failed to read "${packageJsonFilePath}": ${err.message}`);
  }
}

/**
 * @return the `string` field of the given `packageJson` `object` identified
 * by `fieldName`
 */
function readStringField({ fieldName, packageJson }) {
  const fieldValue = packageJson[fieldName];

  if (
    !fieldValue ||
    (typeof fieldValue !== 'string' && !(fieldValue instanceof String))
  ) {
    throw Error(`Package ${fieldName} "${fieldValue}" is invalid`);
  }

  return fieldValue;
}
