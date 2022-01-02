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

/** Re-applies all of the k8s configuration for `@chorro/web`. */
export function deployChorroWeb() {
  const logTag = `[${appsDirectoryName}/${webAppDirectoryName}]`;
  const rootDirectoryPath = resolveProjectRootDirectoryPath();

  const k8sDirectoryPath = asPath(
    rootDirectoryPath,
    appsDirectoryName,
    webAppDirectoryName,
    k8sDirectoryName,
  );

  console.info(logTag, 'Applying k8s deployment configuration...');
  applyK8sConfigurationFile(k8sDeploymentFileName, k8sDirectoryPath);

  console.info(logTag, 'Applying k8s service configuration...');
  applyK8sConfigurationFile(k8sServiceFileName, k8sDirectoryPath);

  console.info(logTag, 'Applying k8s ingress configuration...');
  applyK8sConfigurationFile(k8sIngressFileName, k8sDirectoryPath);
}

/** Invokes `deployChorroWeb` if this file was run directly. */
if (esMain(import.meta)) {
  deployChorroWeb();
}
