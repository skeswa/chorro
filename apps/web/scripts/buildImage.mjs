import { execSync } from 'child_process';
import { join as asPath } from 'path';

import { composeWebDockerImageTag } from './core/dockerUtil.mjs';
import {
  appsDirectoryName,
  dockerfileFileName,
  webAppDirectoryName,
} from './core/fileConstants.mjs';
import { resolveProjectRootDirectoryPath } from './core/fileUtil.mjs';

const rootDirectoryPath = resolveProjectRootDirectoryPath();

const appDirectoryPath = asPath(
  rootDirectoryPath,
  appsDirectoryName,
  webAppDirectoryName,
);

const webDockerImageTag = composeWebDockerImageTag(appDirectoryPath);

console.info(`[${appsDirectoryName}/${webAppDirectoryName}] Building docker image ${webDockerImageTag}...`);

execSync(
  [
    'docker',
    'build',
    ...['-f', asPath(appDirectoryPath, dockerfileFileName)],
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
