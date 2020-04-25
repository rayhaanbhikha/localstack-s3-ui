import React, { useContext, useState } from 'react';
import { withRouter } from 'react-router-dom'
import { ResourceRow } from '../../Components/ResourceRow/ResourceRow'
import { S3Context } from '../../context'
import { ReactComponent as Chevron } from './chevron.svg'

import './styles.css'

const BreadCrums = ({ breadcrums }) => {
    return <div className="breadcrums">
        {breadcrums.map((breadcrum, index) =>
            <>
                {breadcrum}
                &nbsp;
                {index !== breadcrums.length - 1 &&
                    <>
                        <Chevron />
                    &nbsp;
                    </>
                }
            </>
        )}
    </div>
}


const Component = ({ match }) => {

    const data = useContext(S3Context);
    const bucketName = match.params.bucketName;
    const bucketResources = data[bucketName].Resources;
    const initState = {
        resources: bucketResources,
        breadcrums: ["Localstack S3", bucketName]
    }
    const [state, setState] = useState(initState)

    return <div className="buckets-table">
        <div className="table-head-container">
            <div className="table-text">
                <BreadCrums breadcrums={state.breadcrums} />
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
                {
                    state.resources.map((bucketData, index) =>
                        <ResourceRow resource={bucketData} key={`${bucketData.name}-${index}`} setState={setState} />
                    )
                }
            </tbody>
        </table>
    </div>
}

export const Bucket = withRouter(Component)