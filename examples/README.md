# Quick Start
## manuscript.yaml Schema
``
name: <string>  # Required: The name of the task.
specVersion: <string>  # Required: The configuration version.
parallelism: <integer>  # Required: The degree of parallelism.

sources:  # Optional: Defines the data sources.
- name: <string>    # Required: The name of the data source.
  type: <dataset>   # Required: Must be "dataset".
  dataset: <string> # Required: The dataset name.
  filter: [<string>] # Optional: Filter conditions for the dataset.

transforms:  # Optional: Defines data transformation steps.
- name: <string>  # Required: The name of the transform step.
  sql: <string>   # Required: The SQL query used for transformation.

sinks:  # Optional: Defines data sinks (output destinations).
- name: <string>  # Required: The name of the sink.
  type: <string>  # Required: The type of the sink. Supported values: print, filesystem, pg, starrocks.
  from: <string>  # Required: The source of the data, usually a transform step.
  database: [<string>]  # Optional: The database name (for database sinks like pg).
  schema: [<string>]    # Optional: The database schema.
  table: [<string>]     # Optional: The target table name.
  primary_key: [<list>] # Optional: The primary key fields for deduplication or updates.
  config:  # Optional: Additional configuration for the sink.
  host: [<string>]    # Optional: The database host.
  port: [<integer>]   # Optional: The database port.
  username: [<string>] # Optional: The database username.
  password: [<string>] # Optional: The database password.
``
