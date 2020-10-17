import { LogEntry, Log as LogT } from "./../../types/types";
import styles from "./Log.module.css";
import React from "react";

interface Props {
  log: LogT;
};

const Log: React.FC<Props> = ({ log }) => {
  return (
    <div className={styles["log"]}>
    {
      log.map((e: LogEntry, i: number) => {
        return (
          <div className={styles["log-entry"]} key={i}>
            <p>URL: {e.url}</p>
            <p>Stub: <a href={e.stub}>{e.stub}</a></p>
          </div>
        );
      })
    }
    </div>
  );
};

export default Log;
