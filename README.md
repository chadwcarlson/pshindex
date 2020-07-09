# pshindex

> Heavy lifting comes from the [`chadwcarlson/gomeili`](https://github.com/chadwcarlson/gomeili) module.

1. Download Meilisearch and Go modules:

    ```bash
    ./setup.sh
    ```

2. Run the Meilisearch server ([`localhost:7700`](http://localhost:7700/)):

    ```bash
    ./meilisearch
    ```

3. Create a `.env` file in the project root that contains a GitHub access token:

    ```text
    # .env
    GITHUBTOKEN=XXXXXXXX
    ```

4. In another terminal, build the index (config @ `config/index.yaml`) and post to Meilisearch (config @ `config/meili.yaml`):

    ```bash
    ./pshindex

    Collecting documents for Meilisearch...

    Remote Meilisearch index @ https://docs.platform.sh/index.json
     Done. Index written to: output/docs.json

    Platform.sh templates @ https://api.github.com/repos/platformsh/template-builder/contents/templates
     Templates (65 on `master`)
     100% |██████████████████████████████████████████████████████████████████████████████| (65/65, 2 it/s) [30s:0s]
     Done. Index written to: output/templates.json

    Discourse API @ https://community.platform.sh/
     How-to Guides (88 topics)
     100% |███████████████████████████████████████████████████████████████████████████████| (30/30, 5 it/s) [6s:0s]
     100% |███████████████████████████████████████████████████████████████████████████████| (30/30, 5 it/s) [5s:0s]
     100% |███████████████████████████████████████████████████████████████████████████████| (28/28, 5 it/s) [5s:0s]
     Questions & Answers (110 topics)
     100% |██████████████████████████████████████████████████████████████████████████████| (50/50, 5 it/s) [10s:0s]
     100% |██████████████████████████████████████████████████████████████████████████████| (50/50, 5 it/s) [10s:0s]
     100% |███████████████████████████████████████████████████████████████████████████████| (10/10, 5 it/s) [2s:0s]
     Tutorials (14 topics)
     100% |███████████████████████████████████████████████████████████████████████████████| (14/14, 5 it/s) [2s:0s]
     Activity scripts (3 topics)
     100% |█████████████████████████████████████████████████████████████████████████████████| (4/4, 5 it/s) [0s:0s]
     Done. Index written to: output/community.json

    Open API 3.0 specification @ https://api.platform.sh/docs/openapispec-platformsh.json
     Paths (111 paths)
     Tags (22 tags)
     Done. Index written to: output/apidocs.json

    Writing Combined index
     Done. Index written to: output/combined.json

    Posting index to Meilisearch...

    - Current index detecteed
    - Index created.
    - Index name set.
    - Primary key set.
    - Synonyms applied.
    - Stop words updated.
    - Ranking rules updated.
    - Searchable attributes updated.
    - Displayed attributes updated.
    - Distinct attribute updated.
    - Documents added.

    Done.
    ```
