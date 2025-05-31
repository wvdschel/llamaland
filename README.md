# Compute Maestro

Compute maestro is an orchestration tool for hosting different AI and other compute intense services on a single machine.

## Installation

Currently, there are no packaged releases for Compute Maetro. You can install the latest development version of Compute Maestro using `go install`:

    ```
    go install github.com/compute-maestro/compute-maestro@master
    ```

## Planned features

[ ] Automatic loading and unloading of services and models based on requests and available resources
[ ] Logging of queries, prompts and responses for debugging and analysing how the model is being used
[ ] Managing model data & shared model storage between services