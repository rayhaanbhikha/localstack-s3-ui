import React from 'react'
import { ReactComponent as FileIcon } from './file.svg'
import { ReactComponent as FolderIcon } from './folder.svg'
import { ReactComponent as BucketIcon } from './bucket.svg'

import './styles.css'
import { withRouter } from 'react-router-dom'

const Component = ({ history, resource, fetchResources }) => {
  return <tr>
    <td onClick={() => {
      if (resource.type === "File") {
        window.location.href = `http://localhost:8080/page?path=${resource.path}`
      }
      fetchResources(resource)
    }}>
      <div className="resource">
        { resource.type === "Bucket" && <BucketIcon className="icon"/>}
        { resource.type === "Directory" && <FolderIcon className="icon"/>}
        { resource.type === "File" && <FileIcon className="icon"/>}
        <div>
          {resource.name}
        </div>
      </div>
    </td>
  </tr>
}

export const ResourceRow = withRouter(Component)