import React, { useState } from 'react'
import { ReactComponent as FileIcon } from './file.svg'
import { ReactComponent as FolderIcon } from './folder.svg'

import './styles.css'

export const ResourceRow = ({ resource, setState }) => {
  return <tr>
    <td onClick={() => {
      setState(prevState => ({
        ...prevState,
        resources: resource.resources,
        breadcrums: [...prevState.breadcrums, resource.name]
      }))
    }}>
      <div className="resource">
        {resource.type == "File" ? <FileIcon className="icon" /> : <FolderIcon className="icon" />}
        <div>
          {resource.name}
        </div>
      </div>
    </td>
  </tr>
}