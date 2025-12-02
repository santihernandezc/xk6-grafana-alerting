/**
 * **Example k6 extension**
 *
 * @module grafana-alerting
 */
export as namespace grafana_alerting;

/**
 * Generate and optionally create a set of Grafana rules.
 *
 * @param config
 */
export declare function GenerateGroups(config: GenerateConfig): GenerateOutput;

export interface GenerateOutput {
  // TODO: figure out if we can import the open api types directly
  groups: object[];
  // FIXME: change this when we fix the camel case issue in the conversion
  input_config: GenerateConfig;
}

export interface UploadConfig {
  grafanaURL: string;
  username?: string;
  password?: string;
  token?: string;
  orgId?: number;
  folderUIDs: string[];
}

export interface GenerateConfig {
  alertRuleCount: number;
  recordingRuleCount: number;
  queryDatasource: string;
  writeDatasource: string;
  rulesPerGroup: number;
  groupsPerFolder: number;
  seed?: number;
  uploadConfig?: UploadConfig;
}
