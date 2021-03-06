import { constants } from "./actions";
import { constants as phaseConstants } from "../appRoutes/actions";

const stepsDataState = {
  step: {},
  phase: "loading",
  progress: {},
  stepError: "",
  actions: [],
  kustomizeFlow: false,
  isPolling: false,
};

export function stepsData(state = stepsDataState, action) {
  switch (action.type) {
  case constants.RECEIVE_CURRENT_STEP:
    const { currentStep } =  action.payload;
    const isKustomize = currentStep.helmIntro || currentStep.helmValues || currentStep.kustomize;
    return Object.assign({}, state, {
      step: currentStep,
      phase: action.payload.phase,
      progress: action.payload.progress || action.payload.status,
      actions: action.payload.actions,
      kustomizeFlow: isKustomize,
    });
  case phaseConstants.SET_PHASE:
    return Object.assign({}, state, {
      phase: action.payload
    })
  case constants.SET_STEP_ERROR:
    return Object.assign({}, state, {
      stepError: action.payload
    });
  case phaseConstants.SET_PROGRESS:
    return {
      ...state,
      progress: action.payload,
    };
  case phaseConstants.POLLING:
    return {
      ...state,
      isPolling: action.payload,
    };
  default:
    return state;
  }
}
