import React, { useContext } from 'react';
import { withRouter, Link } from 'react-router-dom'
import { S3Context } from '../../context'
import { ReactComponent as FileIcon } from './file.svg'
import { ReactComponent as FolderIcon } from './folder.svg'
import './styles.css'

const Resource = ({ resource }) => {

    const renderFile = () => <div className="resource">
        <FileIcon className="icon" />
        <div>
            {resource.name}
        </div>
    </div>
    const renderDir = () => <div className="resource">
        <FolderIcon className="icon" />
        <div>
            {resource.name}
        </div>
    </div>

    return <tr>
        <td>
            {resource.type == "File" ? renderFile() : renderDir()}
        </td>
    </tr>
}


const Component = ({ match }) => {
    const data = useContext(S3Context);
    const bucketName = match.params.bucketName;
    const bucketResources = data[bucketName].Resources;
    console.log(bucketResources)
    return <div className="buckets-table">
        <div className="table-head-container">
            <div className="table-text">
                <strong className="table-bucket-text">{bucketName}</strong>
            </div>
        </div>
        <table>
            <thead >
                <tr className="table-column-heading">
                    <th className="table-column-heading-text">Name</th>
                </tr>
            </thead>
            <tbody>
                {bucketResources.map((bucketData, index) => <Resource resource={bucketData} key={`${bucketData.name}-${index}`} />)}
            </tbody>
        </table>
    </div>
}

export const Bucket = withRouter(Component)