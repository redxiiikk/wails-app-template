export namespace api {
	
	export class EchoRequest {
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new EchoRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	    }
	}
	export class EchoResponse {
	    message: string;
	
	    static createFrom(source: any = {}) {
	        return new EchoResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.message = source["message"];
	    }
	}
	export class HealthCheckItem {
	    name: string;
	    status: string;
	    errorMessage: string;
	
	    static createFrom(source: any = {}) {
	        return new HealthCheckItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.status = source["status"];
	        this.errorMessage = source["errorMessage"];
	    }
	}
	export class HealthCheckResponse {
	    items: HealthCheckItem[];
	
	    static createFrom(source: any = {}) {
	        return new HealthCheckResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.items = this.convertValues(source["items"], HealthCheckItem);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

