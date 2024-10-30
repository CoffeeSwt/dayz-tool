export namespace common {
	
	export class ServerStartConfig {
	    dayzserverpath: string;
	    clientmods: string;
	    servermods: string;
	    map: string;
	
	    static createFrom(source: any = {}) {
	        return new ServerStartConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.dayzserverpath = source["dayzserverpath"];
	        this.clientmods = source["clientmods"];
	        this.servermods = source["servermods"];
	        this.map = source["map"];
	    }
	}

}

