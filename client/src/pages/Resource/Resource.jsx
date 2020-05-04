import React, { useEffect, useState } from 'react'
import { withRouter } from 'react-router-dom'

const Component = ({ location, history }) => {
  if (!location.state || !location.state.path) {
    history.replace("/s3");
  }

  const { path } = location.state

  const [resource, setResources] = useState({})

  const fetchResources = async () => {
    try {
      console.log("loading bucket data")
      const res = await fetch(`http://localhost:8080/resource?path=${path}`)
      const data = await res.json();
      console.log(data);
      setResources(data)
    } catch (error) {
      console.log(error)
    }
  }

  useEffect(() => {
    fetchResources();
  }, [])

  return (
    <div>
      {resource.name}
    </div>
  )
}

export const Resource = withRouter(Component)