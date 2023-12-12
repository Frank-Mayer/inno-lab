enum LogLevel {
    INFO,
    WARN,
    ERROR,
}

function logLevelToString(level: LogLevel) {
    switch (level) {
        case LogLevel.INFO:
            return "INFO";
        case LogLevel.WARN:
            return "WARN";
        case LogLevel.ERROR:
            return "ERROR";
        default:
            return "UNKNOWN";
    }
}

function build(level: LogLevel, message: string, keyvals: Array<any>) {
    if (keyvals.length % 2 !== 0) {
        throw new Error("log message must have even number of keyvals");
    }

    const sb = new Array<string>();
    sb.push(new Date().toISOString());
    sb.push(" ");
    sb.push(logLevelToString(level));
    sb.push(" ");
    sb.push(message);

    for (let i = 2; i < keyvals.length; i += 2) {
        sb.push(" ");
        const k = keyvals[i];
        if (typeof k !== "string") {
            throw new Error("log message key must be a string");
        }
        sb.push();
        sb.push("=");
        const v = keyvals[i + 1];
        switch (typeof v) {
            case "string":
                sb.push(v);
                break;
            case "number":
                sb.push(v.toString());
                break;
            case "boolean":
                sb.push(v ? "true" : "false");
                break;
            default:
                if ("toString" in v && typeof v.toString === "function") {
                    sb.push(v.toString());
                    break;
                }
                if ("valueOf" in v && typeof v.valueOf === "function") {
                    sb.push(v.valueOf().toString());
                    break;
                }
                sb.push(JSON.stringify(v));
                break;
        }
    }

    return sb.join("");
}

export function error(message: string): void;
export function error(
    message: string,
    ...keyvals: PairArray<string, any>
): void;
export function error(message: string, ...keyvals: Array<any>) {
    console.error(build(LogLevel.ERROR, message, keyvals || []));
}

export function warn(message: string): void;
export function warn(message: string, ...keyvals: PairArray<string, any>): void;
export function warn(message: string, ...keyvals: Array<any>) {
    console.warn(build(LogLevel.WARN, message, keyvals || []));
}

export function info(message: string): void;
export function info(message: string, ...keyvals: PairArray<string, any>): void;
export function info(message: string, ...keyvals: Array<any>) {
    console.info(build(LogLevel.INFO, message, keyvals || []));
}

export function fatal(message: string): void;
export function fatal(
    message: string,
    ...keyvals: PairArray<string, any>
): void;
export function fatal(message: string, ...keyvals: Array<any>) {
    const msg = build(LogLevel.ERROR, message, keyvals || []);
    console.error(msg);
    alert(msg);
    throw new Error(msg);
}

/** Array of even length */
type PairArray<A, B> = [A, B, ...Array<A | B>];
