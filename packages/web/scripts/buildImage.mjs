import { execSync } from 'child_process';
import { join as asPath } from 'path';

import { composeWebDockerImageTag } from './core/dockerUtil.mjs';
import {
  dockerfileFileName,
  packagesDirectoryName,
  webDirectoryName,
} from './core/fileConstants.mjs';
import { resolveProjectRootDirectoryPath } from './core/fileUtil.mjs';

const rootDirectoryPath = resolveProjectRootDirectoryPath();

const packageDirectoryPath = asPath(
  rootDirectoryPath,
  packagesDirectoryName,
  webDirectoryName,
);

const webDockerImageTag = composeWebDockerImageTag(packageDirectoryPath);

console.info(`[packages/web] Building docker image ${webDockerImageTag}...`);

execSync(
  [
    'docker',
    'build',
    ...['-f', asPath(packageDirectoryPath, dockerfileFileName)],
    ...['-t', webDockerImageTag],
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
