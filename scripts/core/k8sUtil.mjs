import { execSync } from 'child_process';
import { readFileSync, writeFileSync } from 'fs';
import { sep } from 'path';
import yaml from 'js-yaml';

import {
  composeDockerImageTag,
  decomposeDockerImageTag,
} from './dockerUtil.mjs';
import { k8sDirectoryName } from './fileConstants.mjs';

/**
 * Regular expression designed to match file paths that look like
 * "anything/k8s/anything.yml".
 */
const K8S_CONFIGURATION_FILE_PATH_PATTERN = RegExp(
  `^.+${sep}${k8sDirectoryName}${sep}[^\.]+\.yml$`,
);

/**
 * Applies the k8s configuration file indicated by
 * `relativeK8sConfigurationFilePath`.
 *
 * @param rootDirectoryPath is the root directory path of the monorepo
 */
export function applyK8sConfigurationFile({
  relativeK8sConfigurationFilePath,
  rootDirectoryPath,
}) {
  const output = execSync(
    ['kubectl', 'apply', '-f', relativeK8sConfigurationFilePath].join(' '),
    {
      cwd: rootDirectoryPath,
      encoding: 'utf-8',
    },
  );

  const sanitizedOutput = output.trim().toLowerCase();
  if (sanitizedOutput.endsWith(' unchanged')) {
    console.log(
      relativeK8sConfigurationFilePath,
      'did not update cluster configuration.',
    );

    return { didConfigure: false };
  } else {
    console.log(
      relativeK8sConfigurationFilePath,
      'updated cluster configuration successfully!',
    );

    return { didConfigure: true };
  }
}

/** @returns `true` if `filePath` refers to a k8s configuration file */
export function isK8sConfigurationFilePath(filePath) {
  return !!filePath.match(K8S_CONFIGURATION_FILE_PATH_PATTERN);
}

/**
 * Reads and parses the k8s deployment configuration file located at the
 * specified path.
 *
 * @param {string} k8sDeploymentFilePath path to the k8s deployment
 *    configuration file to read
 */
export function readK8sDeployment(k8sDeploymentFilePath) {
  try {
    const k8sDeploymentFile = readFileSync(k8sDeploymentFilePath, 'utf8');

    const k8sDeploymentDocument = yaml.load(k8sDeploymentFile);

    const { containers } = k8sDeploymentDocument.spec.template.spec;
    if (!containers || containers.length !== 1) {
      throw new Error('Deployment did not have exactly one container');
    }

    const { image: k8sDeploymentDocumentImageTag } = containers[0] ?? {
      image: undefined,
    };
    if (!k8sDeploymentDocumentImageTag) {
      throw new Error(
        'Docker image tag of the first container in the deployment was missing',
      );
    }

    return {
      decomposedK8sDeploymentDocumentImageTag: decomposeDockerImageTag(
        k8sDeploymentDocumentImageTag,
      ),
      k8sDeploymentDocument,
    };
  } catch (err) {
    throw new Error(
      `Failed to read k8s deployment configuration ` +
        `from "${k8sDeploymentFilePath}": ${err.message}`,
    );
  }
}

/**
 * Updates the document image tag of the supplied `k8sDeploymentDocument` in
 * accordance with `decomposedK8sDeploymentDocumentImageTag`, and then writes it
 * to the specified `k8sDeploymentFilePath`.
 */
export function writeK8sDeployment({
  decomposedK8sDeploymentDocumentImageTag,
  k8sDeploymentDocument,
  k8sDeploymentFilePath,
}) {
  try {
    k8sDeploymentDocument.spec.template.spec.containers[0].image =
      composeDockerImageTag(decomposedK8sDeploymentDocumentImageTag);

    const k8sDeploymentFile = yaml.dump(k8sDeploymentDocument, {
      sortKeys: true, // a-z the keys
    });

    writeFileSync(k8sDeploymentFilePath, k8sDeploymentFile, {
      encoding: 'utf-8',
    });
  } catch (err) {
    console.error(err);
    throw new Error(
      `Failed to write k8s deployment configuration ` +
        `to "${k8sDeploymentFilePath}": ${err.message}`,
    );
  }
}
