import { Echo } from '../wailsjs/go/api/EchoApi';

export default async function echoApi(message: string): Promise<string> {
  const response = await Echo({ message: message });
  return response.message;
}
