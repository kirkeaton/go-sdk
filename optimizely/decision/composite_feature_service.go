/****************************************************************************
 * Copyright 2019, Optimizely, Inc. and contributors                        *
 *                                                                          *
 * Licensed under the Apache License, Version 2.0 (the "License");          *
 * you may not use this file except in compliance with the License.         *
 * You may obtain a copy of the License at                                  *
 *                                                                          *
 *    http://www.apache.org/licenses/LICENSE-2.0                            *
 *                                                                          *
 * Unless required by applicable law or agreed to in writing, software      *
 * distributed under the License is distributed on an "AS IS" BASIS,        *
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. *
 * See the License for the specific language governing permissions and      *
 * limitations under the License.                                           *
 ***************************************************************************/

package decision

import (
	"github.com/optimizely/go-sdk/optimizely/entities"
)

// CompositeFeatureService is the default out-of-the-box feature decision service
type CompositeFeatureService struct {
	experimentDecisionService        ExperimentDecisionService
	rolloutExperimentDecisionService ExperimentDecisionService
}

// NewCompositeFeatureService returns a new instance of the CompositeFeatureService
func NewCompositeFeatureService(experimentDecisionService ExperimentDecisionService) *CompositeFeatureService {
	if experimentDecisionService == nil {
		experimentDecisionService = NewExperimentBucketerService()
	}
	return &CompositeFeatureService{
		experimentDecisionService: experimentDecisionService,
	}
}

// GetDecision returns a decision for the given feature and user context
func (featureService *CompositeFeatureService) GetDecision(decisionContext FeatureDecisionContext, userContext entities.UserContext) (FeatureDecision, error) {
	featureEnabled := false
	feature := decisionContext.Feature

	// Check if user is bucketed in feature experiment
	experimentDecisionContext := ExperimentDecisionContext{
		Experiment: feature.FeatureExperiments[0],
	}

	experimentDecision, err := featureService.experimentDecisionService.GetDecision(experimentDecisionContext, userContext)
	if err != nil {
		// @TODO(mng): handle error here
	}
	featureEnabled = experimentDecision.Variation.FeatureEnabled

	return FeatureDecision{
		FeatureEnabled: featureEnabled,
	}, nil
}