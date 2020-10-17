import React from "react";
import { NextRouter, useRouter } from "next/router";
const { ShortenerClient } = require("@pb/ShortenerServiceClientPb");
const { Short, Url } = require("@pb/shortener_pb");
import getConfig from "next/config";
const { publicRuntimeConfig } = getConfig();

const Redirect: React.FC = () => {
  let [url, setUrl] = React.useState<string>("");
  let [err, setErr] = React.useState<string>("");
  const router: NextRouter = useRouter()
  const { stub } = router.query;

  React.useEffect(() => {
    if (stub === undefined) return;
    const client = new ShortenerClient(publicRuntimeConfig.GRPC_PROXY, null, null);
    let req = new Short();
    req.setStub(stub);
    client.unShorten(req, null, (err: any, resp: typeof Url) => {
      if (err) setErr(err.message);
      else     setUrl(resp.getUrl());
    });
  }, [stub]);

  React.useEffect(() => {
    if (url.length === 0) return;
    let re = /^https?:///;
    if (re.test(url)) location.href = url;
    else location.href = `http://${url}`;
  }, [url]);

  return <h1 style={{textAlign: "center"}}>{err}</h1>
};

export default Redirect;
