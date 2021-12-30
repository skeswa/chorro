import { execSync } from 'child_process';
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

const rootDirectoryPath = resolveProjectRootDirectoryPath();

const k8sDirectoryPath = asPath(
  rootDirectoryPath,
  appsDirectoryName,
  webAppDirectoryName,
  k8sDirectoryName,
);

console.info('[packages/web] Appling k8s deployment configuration...');

execSync(['kubectl', 'apply', '-f', k8sDeploymentFileName].join(' '), {
  cwd: k8sDirectoryPath,
  // Ignore stdin.
  input: 'ignore',
  // Pipe stdout and stderr to the terminal.
  stdio: 'inherit',
});

console.info('[packages/web] Appling k8s service configuration...');

execSync(['kubectl', 'apply', '-f', k8sServiceFileName].join(' '), {
  cwd: k8sDirectoryPath,
  // Ignore stdin.
  input: 'ignore',
  // Pipe stdout and stderr to the terminal.
  stdio: 'inherit',
});

console.info('[packages/web] Appling k8s ingress configuration...');

execSync(['kubectl', 'apply', '-f', k8sIngressFileName].join(' '), {
  cwd: k8sDirectoryPath,
  // Ignore stdin.
  input: 'ignore',
  // Pipe stdout and stderr to the terminal.
  stdio: 'inherit',
});