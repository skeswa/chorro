import { execSync } from 'child_process';
import esMain from 'es-main';
import { join as asPath, relative as asRelativePath } from 'path';

import { composeDockerImageTag } from './core/dockerUtil.mjs';
import {
  appsDirectoryName,
  dockerfileFileName,
  k8sDeploymentFileName,
  k8sDirectoryName,
  webAppDirectoryName,
} from './core/fileConstants.mjs';
import { resolveProjectRootDirectoryPath } from './core/fileUtil.mjs';
import { readK8sDeployment, writeK8sDeployment } from './core/k8sUtil.mjs';
import { readPackageVersion } from './core/nodeUtil.mjs';

/**
 * Prepares `@chorro/web` for deployment within a Kubernetes cluster:
 * -  If package version differs from k8s deployment config, then:
 *    1.  Runs `yarn build` in the root of the monorepo
 *    2.  Builds a new Docker image for the new package version
 *    3.  Pushes the newly build Docker image to the container registry
 *    4.  Updates the k8s deployment config locally (whoeever invoked this
 *        `function` is expected to commit and push the the updated k8s
 *        deployment config)
 * -  If package version **does NOT** differs from k8s deployment config, then:
 *    - Does nothing
 */
export function releaseChorroWeb() {
  const logTag = `[${appsDirectoryName}/${webAppDirectoryName}]`;
  const rootDirectoryPath = resolveProjectRootDirectoryPath();

  const appDirectoryPath = asPath(
    rootDirectoryPath,
    appsDirectoryName,
    webAppDirectoryName,
  );

  const k8sDeploymentFilePath = asPath(
    appDirectoryPath,
    k8sDirectoryName,
    k8sDeploymentFileName,
  );

  const { decomposedK8sDeploymentDocumentImageTag, k8sDeploymentDocument } =
    readK8sDeployment(k8sDeploymentFilePath);

  const { dockerImageVersion: k8sDeploymentDocumentImageVersion } =
    decomposedK8sDeploymentDocumentImageTag;
  const packageVersion = readPackageVersion(appDirectoryPath);

  if (k8sDeploymentDocumentImageVersion === packageVersion) {
    console.info(
      logTag,
      'K8s deployment is in sync with package.json already - ' +
        'exiting early',
    );

    return;
  }

  const updatedDecomposedK8sDeploymentDocumentImageTag = {
    ...decomposedK8sDeploymentDocumentImageTag,
    dockerImageVersion: packageVersion,
  };

  const updatedK8sDeploymentDocumentImageTag = composeDockerImageTag(
    updatedDecomposedK8sDeploymentDocumentImageTag,
  );

  console.info(
    logTag,
    `Building and pushing docker image`,
    `${updatedK8sDeploymentDocumentImageTag}...`,
  );

  execSync(
    [
      'docker',
      'build',
      ...['-f', asPath(appDirectoryPath, dockerfileFileName)],
      ...['-t', updatedK8sDeploymentDocumentImageTag],
      '.',
    ].join(' '),
    {
      cwd: rootDirectoryPath,
      // Ignore stdin.
      input: 'ignore',
      // Pipe stdout and stderr to the terminal.
      stdio: 'inherit',
    },
  );

  execSync(['docker', 'push', updatedK8sDeploymentDocumentImageTag].join(' '), {
    cwd: rootDirectoryPath,
    // Ignore stdin.
    input: 'ignore',
    // Pipe stdout and stderr to the terminal.
    stdio: 'inherit',
  });

  console.info(
    logTag,
    `Updating K8s deployment image tag version from`,
    `${k8sDeploymentDocumentImageVersion} → ${packageVersion}...`,
  );

  writeK8sDeployment({
    decomposedK8sDeploymentDocumentImageTag: {
      ...decomposedK8sDeploymentDocumentImageTag,
      dockerImageVersion: packageVersion,
    },
    k8sDeploymentDocument,
    k8sDeploymentFilePath,
  });

  execSync(
    ['git add', asRelativePath(rootDirectoryPath, k8sDeploymentFilePath)].join(
      ' ',
    ),
    {
      cwd: rootDirectoryPath,
      // Ignore stdin.
      input: 'ignore',
      // Pipe stdout and stderr to the terminal.
      stdio: 'inherit',
    },
  );
}

/** Invokes `releaseChorroWeb` if this file was run directly. */
if (esMain(import.meta)) {
  releaseChorroWeb();
}
