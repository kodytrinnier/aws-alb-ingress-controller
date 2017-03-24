package controller

import "github.com/aws/aws-sdk-go/service/elbv2"

type Listeners []*Listener

func (l Listeners) find(listener *elbv2.Listener) int {
	for p, v := range l {
		if v.Equals(listener) {
			return p
		}
	}
	return -1
}

// // Meant to be called when we delete a targetgroup and just need to lose references to our listeners
// func (l Listeners) purgeTargetGroupArn(a *ALBIngress, arn *string) Listeners {
// 	var listeners Listeners
// 	for _, listener := range l {
// 		// TODO: do we ever have more default actions?
// 		if *listener.CurrentListener.DefaultActions[0].TargetGroupArn != *arn {
// 			listeners = append(listeners, listener)
// 		}
// 	}
// 	return listeners
// }

// func (l Listeners) modify(a *ALBIngress, lb *LoadBalancer) error {
// 	var li Listeners
// 	for _, targetGroup := range lb.TargetGroups {
// 		for _, listener := range lb.Listeners {
// 			// TODO this may bomb if it needs to create a rule that depends on a new listener
// 			// rules := listener.Rules.delete()
// 			rules := listener.Rules.modify(a, listener, targetGroup)
// 			listener.Rules = rules

// 			if listener.DesiredListener == nil {
// 				listener.delete(a)
// 				continue
// 			}
// 			if listener.CurrentListener == nil {
// 				listener.create(a, lb, targetGroup)
// 				continue
// 			}
// 			// rules := listener.Rules.create()
// 			if err := listener.modify(a, lb, targetGroup); err != nil {
// 				return err
// 			}
// 			li = append(li, listener)
// 		}
// 	}
// 	lb.Listeners = li
// 	return nil
// }

// func (l Listeners) delete(a *ALBIngress) error {
// 	errors := false
// 	for _, listener := range l {
// 		if err := listener.delete(a); err != nil {
// 			glog.Infof("%s: Unable to delete listener %s: %s",
// 				a.Name(),
// 				*listener.CurrentListener.ListenerArn,
// 				err)
// 		}
// 	}
// 	if errors {
// 		return fmt.Errorf("There were errors deleting listeners")
// 	}
// 	return nil
// }

func (l Listeners) StripDesiredState() {
	for _, listener := range l {
		listener.DesiredListener = nil
	}
}
