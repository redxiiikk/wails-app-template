import { HealthCheck } from '../wailsjs/go/api/HealthCheckApi';

export interface HealthStatus {
  systemName: string;
  status: string;
  message: string;
}

export default async function healthCheckApi(): Promise<HealthStatus[]> {
  const response = await HealthCheck();

  return response.items.map((item) => ({
    systemName: item.name,
    status: item.status,
    message: item.errorMessage,
  }));
}
