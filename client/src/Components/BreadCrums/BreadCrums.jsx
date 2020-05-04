import React from 'react';
import { ReactComponent as Chevron } from './chevron.svg'

import './styles.css'

export const BreadCrums = ({ breadcrums }) => {
    return <div className="breadcrums">
        {breadcrums.map((breadcrum, index) =>
            <>
                {breadcrum.label}
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