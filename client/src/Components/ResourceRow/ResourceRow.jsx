import React from 'react'
import { ReactComponent as FileIcon } from './file.svg'
import { ReactComponent as FolderIcon } from './folder.svg'

import './styles.css'

export const ResourceRow = ({ resource, setState }) => {
  return <tr>
    <td onClick={() => {
      if (resource.type === "File") {
        setState(prevState => ({
          type: "File",
          resource: resource,
          breadcrums: [...prevState.breadcrums, resource.name]
        }))
        return
      } 
      setState(prevState => ({
        type: "Dir",
        resources: resource.resources,
        breadcrums: [...prevState.breadcrums, resource.name]
      }))
    }}>
      <div className="resource">
        {resource.type === "File" ? <FileIcon className="icon" /> : <FolderIcon className="icon" />}
        <div>
          {resource.name}
        </div>
      </div>
    </td>
  </tr>
}