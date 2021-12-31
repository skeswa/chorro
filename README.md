# chorro

Chore scheduling made easy

## Development

Chorro is a
[monorepo](https://en.wikipedia.org/wiki/Monorepo#:~:text=In%20version%20control%20systems%2C%20a,stored%20in%20the%20same%20repository.&text=Many%20attempts%20have%20been%20made,other%2C%20newer%20forms%20of%20monorepos.)
managed by [turborepo](https://turborepo.org/).

```
chorro/
├─ .changeset/      # changeset files live in here
├─ apps/            # deployable apps live in here
│  ├─ foo/
│  ├─ bar/
├─ packages/        # reusable packages live in here
│  ├─ foo/
│  ├─ bar/
├─ scripts/         # node scripts used to manage the repo live in here
```

### Contributing

Chorro has only one main `git` branch - `main`. The `main` branch only accepts changes made to it via
[pull request](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/about-pull-requests). Chorro uses [changesets](https://github.com/changesets/changesets) to provide semantic change management for each of the app.

Every pull request (PR) should be accompanied by the addition of a new changeset - **PRs cannot be merged without including a new changeset**
([this functionality is made possible by the changeset bot](https://github.com/apps/changeset-bot)).

Adding a new changeset is easy:

```bash
$ yarn changeset
```
