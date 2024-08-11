import { MigrateHistory as MigrateHistoryApi } from '../wailsjs/go/api/MigrateHistoryApi';

export interface MigrateHistory {
  id: number;
  key: string;
  hash: string;
  createdAt: Date;
}

export default async function queryAllMigrateHistoryApi(): Promise<
  MigrateHistory[]
> {
  const response = await MigrateHistoryApi();

  return response.map((item) => ({
    id: item.id,
    key: item.key,
    hash: item.hash,
    createdAt: new Date(item.createdAt),
  }));
}
