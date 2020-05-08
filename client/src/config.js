import { joinPath } from "./utils";

const getConfig = () => {
  const host = process.env.NODE_ENV === "production" ? window.location.origin : process.env.REACT_APP_HOST
  const path = process.env.REACT_APP_API_PATH
  return {
    host,
    path,
    apiUrl: joinPath(host, path)
  }
}

export const config = getConfig();