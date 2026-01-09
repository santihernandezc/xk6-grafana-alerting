export function envOrDefault(envVarName, def) {
  const envVar = __ENV[envVarName];
  return envVar ?? def;
}

// ensureConfig builds a configuration for the API using environment variables and defaults.
export function ensureConfig() {
  return {
    url: envOrDefault("GRAFANA_URL", "http://localhost:3000"),
    username: envOrDefault("GRAFANA_ADMIN_USER", "admin"),
    password: envOrDefault("GRAFANA_ADMIN_PASSWORD", "admin"),
    token: envOrDefault("GRAFANA_API_TOKEN", ""),
  };
}
