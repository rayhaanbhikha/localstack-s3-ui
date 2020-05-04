import React from 'react';
import { ReactComponent as Chevron } from './chevron.svg'

import './styles.css'

export const BreadCrums = ({ fetchResources, breadcrums }) => {
    return <div className="breadcrums">
        {breadcrums.map((breadcrum, index) =>
            <>
                <div onClick={() => fetchResources(breadcrum.url)}>
                    {breadcrum.label}
                </div>
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