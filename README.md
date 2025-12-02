# xk6-grafana-alerting

**Grafana Alerting k6 extension**

This extension is experimental! It provides functionality to generate Grafana Alerting rules and recording rules for load testing purposes.


```javascript file=script.js
import { GenerateGroups } from "k6/x/grafana-alerting";

export default function () {
  let output = GenerateGroups({
    alertRuleCount: 20,
    recordingRuleCount: 30,
    queryDatasource: "query_ds_id",
    writeDatasource: "write_ds_id",
    rulesPerGroup: 4,
    groupsPerFolder: 5,
  });
  let groups = output.groups;
  let inputConfig = output.input_config;
  console.log(output);
  console.log(groups);
  console.log(groups.length);
  console.log(groups[0].title);
  console.log(groups[0].rules.length);
  console.log(groups[0].rules[0]);
  console.log(groups[0].rules[0].title);
  console.log(inputConfig);
}
```

## Quick start

Run the test script. The repository's root directory includes a `script.js` file. When developing k6 extensions, use the `xk6 run` command instead of `k6 run` to execute your scripts.

    ```shell
    xk6 run script.js
    ```

## Development environment

While using a GitHub codespace in the browser is a good starting point, you can also set up a local development environment for a better developer experience.

To create a local development environment, you need an IDE that supports [Development Containers](https://containers.dev/). [Visual Studio Code](https://code.visualstudio.com/) supports Development Containers after installing the [Dev Containers extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers).

1. Visual Studio Code will detect the [development container](https://containers.dev/) configuration and show a pop-up to open the project in a dev container. Accept the prompt and the project opens in the dev container, and the container image is rebuilt if necessary.

1. Run the test script. The repository's root directory includes a `script.js` file. When developing k6 extensions, use the `xk6 run` command instead of `k6 run` to execute your scripts.

    ```shell
    xk6 run script.js
    ```

## Download

Building a custom k6 binary with the `xk6-grafana-alerting` extension is necessary for its use. You can download pre-built k6 binaries from the [Releases page](https://github.com/grafana/xk6-grafana-alerting/releases/).

## Build

Use the [xk6](https://github.com/grafana/xk6) tool to build a custom k6 binary with the `xk6-grafana-alerting` extension. Refer to the [xk6 documentation](https://github.com/grafana/xk6) for more information.

## Contribute

If you wish to contribute to this project, please start by reading the [Contributing Guidelines](CONTRIBUTING.md).
