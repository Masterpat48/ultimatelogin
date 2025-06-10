export namespace main {
	
	export class LoginResult {
	    success: boolean;
	    name: string;
	    permission: string;
	
	    static createFrom(source: any = {}) {
	        return new LoginResult(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.success = source["success"];
	        this.name = source["name"];
	        this.permission = source["permission"];
	    }
	}

}

