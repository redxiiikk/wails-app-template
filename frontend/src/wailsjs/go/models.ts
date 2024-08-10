export namespace api {
  export class EchoRequest {
    message: string;

    static createFrom(source: any = {}) {
      return new EchoRequest(source);
    }

    constructor(source: any = {}) {
      if ('string' === typeof source) source = JSON.parse(source);
      this.message = source['message'];
    }
  }
  export class EchoResponse {
    message: string;

    static createFrom(source: any = {}) {
      return new EchoResponse(source);
    }

    constructor(source: any = {}) {
      if ('string' === typeof source) source = JSON.parse(source);
      this.message = source['message'];
    }
  }
}
