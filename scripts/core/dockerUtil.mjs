import * as dockerConstants from './dockerConstants.mjs';

/**
 * Matches, and captures parts of, docker image tag URLs
 * (e.g. ghcr.io/ayy/lmao:latest).
 *
 * Capture groups:
 *
 * | index | description               |
 * |-------|---------------------------|
 * | 0     | container registry domain |
 * | 1     | docker image owner        |
 * | 2     | docker image name         |
 * | 3     | docker image version      |
 */
const DOCKER_IMAGE_TAG_PATTERN = /^([^\/]+)\/([^\/]+)\/([^:]+):(.+)$/;

/**
 * Creates the Docker image tag for the specified app.
 *
 * @param {string} containerRegistryDomain domain of the container registry used to host the docker image
 * @param {string} dockerImageName name of app being tagged
 * @param {string} dockerImageOwner owner of app being tagged
 * @param {string} dockerImageVersion version of app being tagged
 */
export function composeDockerImageTag({
  containerRegistryDomain = dockerConstants.containerRegistryDomain,
  dockerImageName,
  dockerImageOwner,
  dockerImageVersion,
}) {
  const composedDockerImageTag =
    `${containerRegistryDomain}/${dockerImageOwner}/` +
    `${dockerImageName}:${dockerImageVersion}`;

  if (
    !containerRegistryDomain ||
    !dockerImageName ||
    !dockerImageOwner ||
    !dockerImageVersion
  ) {
    throw Error(
      `Docker image tag components don't look quite right: ` +
        `"${composedDockerImageTag}"`,
    );
  }

  return composedDockerImageTag;
}

/**
 * Parses the provided Docker image tag string, returning its components.
 *
 * @param {string} dockerImageTag `string` that looks like
 *    `"ghcr.io/skeswa/chorro-web:0.0.1"`
 * @returns `object` version of `dockerImageTag`
 */
export function decomposeDockerImageTag(dockerImageTag) {
  const matches = DOCKER_IMAGE_TAG_PATTERN.exec(dockerImageTag);
  if (!matches) {
    throw Error(`"${dockerImageTag}" is not a valid docker image tag`);
  }

  const [
    ,
    containerRegistryDomain,
    dockerImageOwner,
    dockerImageName,
    dockerImageVersion,
  ] = matches;

  return {
    containerRegistryDomain,
    dockerImageName,
    dockerImageOwner,
    dockerImageVersion,
  };
}
