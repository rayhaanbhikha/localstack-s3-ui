const joinPathReducer = (path, newPath) => {
  const pathLen = path.length
  const pathLastChar = path[pathLen - 1]
  const newPathFirstChar = newPath[0]

  if(pathLastChar === "/" && newPathFirstChar === "/") {
    return path.substring(0,pathLen-1) + newPath
  } else if (pathLastChar !== "/" && newPathFirstChar !== "/") {
    return path + "/" + newPath
  } else {
    return path + newPath
  }
}

export const joinPath = (currentPath, ...paths) => {
  console.log(currentPath, paths)
  return paths.reduce(joinPathReducer, currentPath)
}

// console.log(joinPath("http://localhost:8080/api/resource/", "/static-resources"))
// console.log(joinPath("http://localhost:8080/api/resource", "/static-resources/", "/hello wolrd/"))
// console.log(joinPath("http://localhost:8080/api/resource/", "static-resources"))
// console.log(joinPath("http://localhost:8080/api/resource", "static-resources", "/hello world"))

export const generateBreadCrums = (str) => {
  const splitStr = str.split("/")
  console.log(splitStr.slice(1))
  let accString = "/"
  return splitStr.slice(1).map(pathName => {
    accString += joinPath(pathName, "/")

    return {
      path: accString,
      name: pathName
    }
  })
}