import React from 'react'

export const Resource = ({ resource }) => {
  const data = "PCFET0NUWVBFIGh0bWw+CjxodG1sIGxhbmc9ImVuIj4KPGhlYWQ+CiAgICA8bWV0YSBjaGFyc2V0PSJVVEYtOCI+CiAgICA8bWV0YSBuYW1lPSJ2aWV3cG9ydCIgY29udGVudD0id2lkdGg9ZGV2aWNlLXdpZHRoLCBpbml0aWFsLXNjYWxlPTEuMCI+CiAgICA8dGl0bGU+RG9jdW1lbnQ8L3RpdGxlPgo8L2hlYWQ+Cjxib2R5PgogICAgSGVsbG8gd29ybGQgdGhpcyBpcyBhbm90aGVyIHRlc3QKPC9ib2R5Pgo8L2h0bWw+"
  return (
    <div>
      name: {resource.name}
      <div>
        {resource.data}
      </div>
      <a href={`http://localhost:8080/echo?data=${data}`}></a>
    </div>
  )
}
