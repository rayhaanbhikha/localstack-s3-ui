import React from 'react';
import { withRouter } from 'react-router-dom'


const Component = ({ match }) => {
    const bucketName = match.params.bucketName;

    return <div>{bucketName}</div>
}

export const Bucket = withRouter(Component)