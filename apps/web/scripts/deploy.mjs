import esMain from 'es-main';
import { join as asPath } from 'path';

import {
  appsDirectoryName,
  k8sDeploymentFileName,
  k8sDirectoryName,
  k8sIngressFileName,
  k8sServiceFileName,
  webAppDirectoryName,
} from './core/fileConstants.mjs';
import { resolveProjectRootDirectoryPath } from './core/fileUtil.mjs';
import { applyK8sConfigurationFile } from './core/k8sUtil.mjs';
import { readPackageMetadata } from './core/nodeUtil.mjs';

/** Re-applies all of the k8s configuration for `@chorro/web`. */
export function deployChorroWeb() {
  const logTag = `[${appsDirectoryName}/${webAppDirectoryName}]`;
  const rootDirectoryPath = resolveProjectRootDirectoryPath();

  const appDirectoryPath = asPath(
    rootDirectoryPath,
    appsDirectoryName,
    webAppDirectoryName,
  );

  const k8sDirectoryPath = asPath(appDirectoryPath, k8sDirectoryName);

  console.info(logTag, 'Applying k8s deployment configuration...');

  const k8sDeploymentUpdate = applyK8sConfigurationFile({
    relativeK8sConfigurationFilePath: k8sDeploymentFileName,
    rootDirectoryPath: k8sDirectoryPath,
  });
  if (k8sDeploymentUpdate.didConfigure) {
    const { name, version } = readPackageMetadata(appDirectoryPath);

    // NOTE: this is a hack !skeswa added so that
    // `https://github.com/changesets/action` automatically recognizes that this
    // package was "published" successfully. Here, we mimic the output of a
    // successful invocation of `changesets publish`.
    console.log(`New tag: ${name}@${version}`);
  }

  console.info(logTag, 'Applying k8s service configuration...');
  applyK8sConfigurationFile({
    relativeK8sConfigurationFilePath: k8sServiceFileName,
    rootDirectoryPath: k8sDirectoryPath,
  });

  console.info(logTag, 'Applying k8s ingress configuration...');
  applyK8sConfigurationFile({
    relativeK8sConfigurationFilePath: k8sIngressFileName,
    rootDirectoryPath: k8sDirectoryPath,
  });
}

/** Invokes `deployChorroWeb` if this file was run directly. */
if (esMain(import.meta)) {
  deployChorroWeb();
}
