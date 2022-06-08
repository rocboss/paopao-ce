import type { Plugin } from "vite"
import { resolve } from "path"
import { readFileSync, writeFileSync } from "fs"

function getFormatTime() {
    let date = new Date();
    let year: string | number = date.getFullYear();
    let month: string | number = date.getMonth() + 1;
    let day: string | number = date.getDate();
    let h: string | number = date.getHours();
    let m: string | number = date.getMinutes();
    let s: string | number = date.getSeconds();
    if (month < 10) month = "0"+month;
    if (day < 10) day = "0"+day;
    if (h < 10) h = "0"+h;
    if (m < 10) m = "0"+m;
    if (s < 10) s = "0"+s;
    return [year, month, day].join("-") + " " + [h, m, s].join(":");
}

/**
 * 自动修改版本号
 * @returns 
 */
export function changePackageVersion(): Plugin {
    if (process.env.NODE_ENV === "production") {
        let packagePath = resolve(__dirname, "./info.json");
        let packageStr = readFileSync(packagePath, {
            encoding: "utf-8"
        });
        let packageJson = JSON.parse(packageStr);
        let version = packageJson.version;
        let versionArr: (string | number)[] = version.split(".");
        let lastVersionNum = +versionArr[versionArr.length - 1];
        lastVersionNum++;
        versionArr.splice(versionArr.length - 1, 1, lastVersionNum);
        packageJson.version = versionArr.join(".");
        packageJson.buildTime = getFormatTime();
        writeFileSync(packagePath, JSON.stringify(packageJson, null, "\t"), {
            encoding: "utf-8"
        })
    }
    return {
        name: "changePackageVersion"
    }
}