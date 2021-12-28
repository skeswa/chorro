import { execSync } from 'child_process';
import { join as asPath } from 'path';

import {
  k8sDeploymentFileName,
  k8sDirectoryName,
  k8sServiceFileName,
  packagesDirectoryName,
  webDirectoryName,
} from './core/fileConstants.mjs';
import { resolveProjectRootDirectoryPath } from './core/fileUtil.mjs';

const rootDirectoryPath = resolveProjectRootDirectoryPath();

const k8sDirectoryPath = asPath(
  rootDirectoryPath,
  packagesDirectoryName,
  webDirectoryName,
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
