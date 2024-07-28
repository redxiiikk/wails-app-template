import {Echo} from "../wailsjs/go/api/EchoApi";

export default async function echoApi(message: string): Promise<string> {
    let response = await Echo({message: message});
    return response.message;
}
