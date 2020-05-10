import React from 'react'
import { withRouter } from 'react-router-dom'

import { ReactComponent as FileIcon } from './file.svg'
import { ReactComponent as FolderIcon } from './folder.svg'
import { ReactComponent as BucketIcon } from './bucket.svg'

import './styles.css'
import { joinPath } from '../../utils';
import { config } from '../../config';

const Component = ({ resource, fetchResources }) => {

  const onClickHandler = () => {
    if (resource.type === "File") {
      window.location.href = joinPath(config.host, resource.resourcePath)
      return
    }
    fetchResources(resource.path)
  }

  return <tr className="resource-row" onClick={onClickHandler}>
    <td >
      <div className="resource">
        {resource.type === "Bucket" && <BucketIcon className="icon" />}
        {resource.type === "Directory" && <FolderIcon className="icon" />}
        {resource.type === "File" && <FileIcon className="icon" />}
        <div>
          {resource.name}
        </div>
      </div>
    </td>
  </tr>
}

export const ResourceRow = withRouter(Component)