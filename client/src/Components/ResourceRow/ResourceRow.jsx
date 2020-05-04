import React from 'react'
import { ReactComponent as FileIcon } from './file.svg'
import { ReactComponent as FolderIcon } from './folder.svg'

import './styles.css'
import { withRouter } from 'react-router-dom'

const Component = ({ history, resource, setPath }) => {
  return <tr>
    <td onClick={() => {
      if (resource.type === "File") {
        history.replace(`/s3/resource/${resource.name}`, { path: resource.path })
      }
      setPath(resource.path)
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

export const ResourceRow = withRouter(Component)