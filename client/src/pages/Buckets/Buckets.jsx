import React, { useContext } from 'react'
import { Link } from 'react-router-dom'
import { S3Context } from '../../context'
import './styles.css'

const BucketRow = ({ bucketName }) => <tr>
    <td>
        <div>
            <Link to={`/s3/${bucketName}`}>
                {bucketName}
            </Link>
        </div>
    </td>
</tr>

export const Buckets = () => {
    const data = useContext(S3Context)
    const bucketNames = Object.entries(data).map(([bucketName]) => bucketName)

    return (
        <div className="buckets-table">
            <div className="table-head-container">
                <div className="table-text">
                    <strong className="table-bucket-text">Buckets</strong>
                    &nbsp;&nbsp;
                    <strong className="table-bucket-nums">({bucketNames.length})</strong>
                </div>
            </div>
            <table>
                <thead >
                    <tr className="table-column-heading">
                        <th className="table-column-heading-text">Name</th>
                    </tr>
                </thead>
                <tbody>
                    {bucketNames.map((bucketName, index) => <BucketRow key={`bucketName-${index}`} bucketName={bucketName} />)}
                </tbody>
            </table>
        </div>
    )
}
