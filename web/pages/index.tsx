import React from "react";
import styles from "./index.module.css";
import { Log as LogT } from "./../src/types/types";
import Log from "@components/log/Log";
const { ShortenerClient } = require("@pb/ShortenerServiceClientPb");
const { ShortenReq, Short } = require("@pb/shortener_pb");
import getConfig from "next/config";
const { publicRuntimeConfig } = getConfig();


const Home: React.FC = () => {
  let [url, setUrl] = React.useState<string>("");
  let [err, setErr] = React.useState<string>("");
  let [log, setLog] = React.useState<LogT>([]);

  const shorten = (url: string, reqStub: string) => {
    let client = new ShortenerClient(publicRuntimeConfig.GRPC_PROXY, null, null),
    req = new ShortenReq();
    req.setUrl(url);
    req.setRequestedstub(reqStub);
    client.shorten(req, null, (err: any, resp: typeof Short) => {
      console.log("callback", err, resp);
      if (err) { setErr(err.message); return; }
      setLog([ ...log, { url: url, stub: resp.getStub() } ]);
    });
  }

  return (
    <div className={styles["app"]}>
      <h1>URL Shortener</h1>
      <div className={styles["app-main"]}>
        <input value={url}
               onChange={(e) => setUrl(e.target.value)}
               className={styles["app-field"]}
               placeholder="Enter URL to shorten"
        />
        <button className={styles["app-button"]}
                onClick={() => { shorten(url, "") }}
        >
          Shorten
        </button>
      </div>
      { err !== "" && <p className={styles["app-error"]}>Error: {err}</p> }
      <Log log={log} />
    </div>
  );
}

export default Home;
